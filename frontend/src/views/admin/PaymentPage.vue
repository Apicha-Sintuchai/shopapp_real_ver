<template>
  <div>
    <div class="flex flex-col justify-center p-20 gap-y-6">
      <div class="flex justify-end items-center"></div>
      <table class="w-full border-collapse overflow-x:auto">
        <thead>
          <tr>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">ปี-เดือน-วัน</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">โต๊ะที่</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">กี่รายการ</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">รวมเป็นเงิน</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">ชำระเงิน</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">รูปแบบชำระเงิน</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="items in filteredItems" :key="items.ID">
            <td class="p-4 text-left border-b-2 border-gray-200">
              {{ items.dateOnly }}
            </td>
            <td class="p-4 text-left border-b-2 border-gray-200">
              {{ items.customer_at_table }}
            </td>
            <td class="p-4 text-left border-b-2 border-gray-200">
              {{ items.orders.length }}
            </td>
            <td class="p-4 text-left border-b-2 border-gray-200">
              {{ items.totalprice }}
            </td>
            <td class="p-4 text-left border-b-2 border-gray-200">
              {{ items.status }}
            </td>

            <td class="flex gap-x-3 p-4 text-left border-b-2 border-gray-200">
              <button
                class="bg-blue-500 text-white px-4 py-2 rounded-lg"
                @click="confirmpaymentcash(items.ID)"
              >
                ชำระเงินสด
              </button>
              <button
                class="bg-red-500 text-white px-4 py-2 rounded-lg"
                @click="confirmPromptPay(items.ID, items.totalprice)"
              >
                ชำระเงินโอน
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import FetchApi from '@/fetch/FetchApi'
import Swal from 'sweetalert2'
import { onMounted, ref, computed } from 'vue'
import generatePayload from 'promptpay-qr'
import QRCode from 'qrcode'
import { jsPDF } from 'jspdf'
import 'jspdf-autotable'

export default {
  setup() {
    const item = ref([])
    const day = ref(null)

    const Getmoney = async () => {
      const response = await FetchApi.Getdata_unpayment()
      item.value = response.data
      item.value.forEach((valve) => {
        const fullDateTime = valve.CreatedAt.split('T')
        valve.dateOnly = fullDateTime[0]
      })
    }

    const filteredItems = computed(() => {
      return item.value.filter((valve) => valve.status === 'ยังไม่ชำระเงิน')
    })

    const generateReceiptPDF = async (id, amount, paymentMethod) => {
      const doc = new jsPDF({
        orientation: 'portrait',
        unit: 'mm',
        format: 'a4',
      })

      doc.setFontSize(24)
      doc.setTextColor(41, 128, 185)
      doc.text('Restaurant Name', doc.internal.pageSize.width / 2, 20, { align: 'center' })

      doc.setFontSize(18)
      doc.setTextColor(0, 0, 0)
      doc.text('Payment Receipt', doc.internal.pageSize.width / 2, 30, { align: 'center' })

      doc.setDrawColor(200, 200, 200)
      doc.line(20, 35, doc.internal.pageSize.width - 20, 35)

      const selectedOrder = item.value.find((order) => order.ID === id)

      doc.setFontSize(12)
      doc.setTextColor(80, 80, 80)
      doc.text('Order Details:', 20, 45)
      doc.text(`Table: ${selectedOrder.customer_at_table}`, 20, 52)
      doc.text(`Order ID: ${id}`, 20, 59)
      doc.text(`Date: ${selectedOrder.dateOnly}`, doc.internal.pageSize.width - 60, 52)
      doc.text(`Total: ${amount} THB`, doc.internal.pageSize.width - 60, 59)
      doc.text(`Payment Method: ${paymentMethod}`, 20, 66)

      const orderItems = selectedOrder.orders.map((order, index) => [
        index + 1,
        order.name_menu,
        order.price,
      ])

      doc.autoTable({
        startY: 75,
        head: [['No.', 'Item Name', 'Price (THB)']],
        body: orderItems,
        theme: 'grid',
        headStyles: {
          fillColor: [41, 128, 185],
          textColor: 255,
          fontSize: 12,
          halign: 'center',
          fontStyle: 'bold',
        },
        bodyStyles: {
          fontSize: 10,
          halign: 'center',
          textColor: [50, 50, 50],
        },
        columnStyles: {
          0: { cellWidth: 20 },
          1: { cellWidth: 100 },
          2: { cellWidth: 40 },
        },
        margin: { left: 20, right: 20 },
        alternateRowStyles: {
          fillColor: [245, 245, 245],
        },
      })

      const footerY = doc.previousAutoTable.finalY + 15
      doc.setFontSize(10)
      doc.setTextColor(128, 128, 128)
      doc.text('Thank you for your business!', doc.internal.pageSize.width / 2, footerY, {
        align: 'center',
      })

      doc.save(`Receipt-${id}.pdf`)
    }

    const confirmpaymentcash = async (id) => {
      await Swal.fire({
        icon: 'warning',
        title: 'หลังจากคุณกดยืนยัน ระบบจะเปลี่ยนสถานะเป็นชำระเงินและข้อมูลชุดนี้จะหายไป ยืนยันหรือไม่?',
        showCancelButton: true,
        confirmButtonText: 'ยืนยัน',
      }).then(async (result) => {
        if (result.isConfirmed) {
          await FetchApi.updatepayment(id)
          const selectedOrder = item.value.find((order) => order.ID === id)
          if (selectedOrder) {
            generateReceiptPDF(id, selectedOrder.totalprice, 'Cash')
            Getmoney()
          }
        }
      })
    }

    const confirmPromptPay = async (id, amount) => {
      const { value } = await Swal.fire({
        icon: 'warning',
        title: 'กรุณาใส่ promptpay ถ้าคุณกดยืนยัน ระบบจะเปลี่ยนสถานะเป็นชำระเงินทันที',
        input: 'number',
        inputPlaceholder: 'เลข promptpay',
        showCancelButton: true,
        showConfirmButton: true,
      })

      if (value) {
        const parseNumber = parseInt(amount)
        const payload = await generatePayload(value, { amount: parseNumber })
        const qrCodeDataUrl = await QRCode.toDataURL(payload)

        await FetchApi.updatepayment(id)
        generateReceiptPDF(id, amount, 'PromptPay')

        const doc = new jsPDF({
          orientation: 'portrait',
          unit: 'mm',
          format: 'a4',
        })

        doc.setFontSize(24)
        doc.setTextColor(41, 128, 185)
        doc.text('Restaurant Name', doc.internal.pageSize.width / 2, 20, { align: 'center' })

        doc.setFontSize(18)
        doc.setTextColor(0, 0, 0)
        doc.text('Payment Receipt', doc.internal.pageSize.width / 2, 30, { align: 'center' })

        doc.setDrawColor(200, 200, 200)
        doc.line(20, 35, doc.internal.pageSize.width - 20, 35)

        const finalY = 75
        doc.setFontSize(14)
        doc.text('Scan to Pay via PromptPay', doc.internal.pageSize.width / 2, finalY, {
          align: 'center',
        })

        const qrSize = 50
        const qrX = (doc.internal.pageSize.width - qrSize) / 2
        doc.addImage(qrCodeDataUrl, 'PNG', qrX, finalY + 5, qrSize, qrSize)

        const footerY = finalY + qrSize + 15
        doc.setFontSize(10)
        doc.setTextColor(128, 128, 128)
        doc.text('Thank you for your business!', doc.internal.pageSize.width / 2, footerY, {
          align: 'center',
        })

        doc.save(`Receipt-${id}.pdf`)
      }
    }

    onMounted(() => {
      Getmoney()
    })

    return {
      item,
      day,
      filteredItems,
      confirmpaymentcash,
      confirmPromptPay,
    }
  },
}
</script>