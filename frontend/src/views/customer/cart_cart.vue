<template>
  <div class="mx-32 my-28">
    <h1 class="text-3xl font-bold mb-8">ลูกค้า โต๊ะ {{ orderDetails.customer_at_table }}</h1>
    <table class="min-w-full border-collapse shadow-xl rounded-xl">
      <thead class="bg-gray-100">
        <tr>
          <th class="p-4 text-left">รูป</th>
          <th class="p-4 text-left">ชื่อเมนู</th>
          <th class="p-4 text-left">ราคา</th>
          <th class="p-4 text-left">เพิ่มเติม</th>
          <th class="p-4 text-left">ลบ</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in orderDetails.order"
          :key="item.name_menu"
          class="border-b hover:bg-gray-50"
        >
          <td class="p-4">
            <img
              :src="`http://localhost:8080/Picture/${item.image}`"
              :alt="item.name_menu"
              class="w-20 h-20 rounded-lg object-cover"
            />
          </td>
          <td class="p-4">{{ item.name_menu }}</td>
          <td class="p-4">{{ item.price }} บาท</td>
          <td class="p-4">{{ item.option || '-' }}</td>
          <td class="p-4">
            <button
              @click="deleteordermenu(item)"
              class="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 transition-colors duration-200 flex items-center gap-2"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
                  clip-rule="evenodd"
                />
              </svg>
              ลบ
            </button>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="flex justify-end items-center mt-4">
      <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" @click="confirmorderbuttom">
        ยืนยัน
      </button>
    </div>
  </div>
</template>

<script setup>
import { useUserOrder } from '@/stores/counter'
import { computed } from 'vue'

const userOrder = useUserOrder()
const orderDetails = computed(() => userOrder.orderDetails)


const deleteordermenu = (items) => {
  userOrder.removeOrderItem(items.name_menu)
}

const confirmorderbuttom= async () => {
    await userOrder.sendorder()
} 

</script>
