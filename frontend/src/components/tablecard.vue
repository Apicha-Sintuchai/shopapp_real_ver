<template>

  <div class="mx-20 my-8">
    <div class="flex justify-end items-center">
      <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" @click="tooglesweetaleart">
        Button
      </button>
    </div>
  </div>
  <div class="mx-36 my-48">
    <div class="relative grid grid-cols-4">
      <div v-for="item in title" :key="item.id">
        <div class="max-w-72 bg-gray-300 rounded-xl p-4 shadow-lg">
          <div class="flex flex-col justify-center items-center gap-y-12">
            <img src="/qr-code.png" alt="" class="w-24 h-24" />
            <h1 @click="onclickbyid(item.ID)">{{ item.ID }}</h1>
            <h1 class="font-bold text-3xl">{{ item.table_number }}</h1>
            <div class="flex justify-between items-center gap-8">
              <button @click="printQRCode(item.table_number)" class="bg-blue-500 text-white px-4 py-2 rounded">
                ปริ้น
              </button>
              <button class="bg-yellow-500 text-white px-4 py-2 rounded">
                แก้ไข
              </button>
              <button class="bg-red-500 text-white px-4 py-2 rounded ">
                ลบ
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, reactive } from "vue";
import QRCode from "qrcode";
import FetchApi from "@/fetch/FetchApi.vue";
import Swal from 'sweetalert2'

export default {
  name: "TableCard",
  setup() {
    // State Variables
    const modal = ref(false);
    const title = ref([]);
    const formdata = reactive({
      table_number: ''
    })
    // Methods
    const toggle = () => {
      modal.value = !modal.value;
    };

    const onclickbyid = (id) => {
      console.log(id);
    };

    const printQRCode = async (tableNumber) => {
      try {
        const qrCodeDataUrl = await QRCode.toDataURL(tableNumber);
        const link = document.createElement("a");
        link.href = qrCodeDataUrl;
        link.download = `QRCode-${tableNumber}.png`;
        link.click();
      } catch (err) {
        console.error("Error generating QR code:", err);
      }
    };

    const tooglesweetaleart = () => {
      Swal.fire({
        title: 'เพิ่มโต๊ะ',
        html: `
      <input id="table-input" class="border-4 w-full" placeholder="ใส่ชื่อโต๊ะ"/>
    `,
        showCancelButton: true,
        confirmButtonText: 'ยืนยัน',
        cancelButtonText: 'ยกเลิก',
        preConfirm: () => {
          const inputValue = document.getElementById('table-input').value;
          if (!inputValue) {
            Swal.showValidationMessage('กรุณากรอกชื่อโต๊ะ');
          }
          return inputValue;
        },
      }).then((result) => {
        if (result.isConfirmed) {
          formdata.table_number = result.value;
          console.log('เพิ่มโต๊ะ:', formdata.table_number);
          // ทำการบันทึกหรืออัปเดตข้อมูลตามต้องการ
        }
      });
    };




    onMounted(async () => {
      const res = await FetchApi.methods.GetTable();
      title.value = res.data.data;
    })

    return {
      modal,
      title,
      toggle,
      onclickbyid,
      printQRCode,
      tooglesweetaleart,
    };
  },
};
</script>

<style scoped>
/* Add your styles here if needed */
</style>
