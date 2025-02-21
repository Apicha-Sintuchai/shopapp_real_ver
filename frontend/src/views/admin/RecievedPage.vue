<template>
  <div>
    <div class="flex flex-col justify-center p-20 gap-y-6">
      <div class="flex justify-end items-center">
        <div class="w-full max-w-sm min-w-[200px]">
          <input
            v-model.lazy="day"
            class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
            type="date"
          />
        </div>
      </div>
      <table class="w-full border-collapse">
        <thead>
          <tr>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">ปี-เดือน-วัน</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">โต๊ะที่</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">กี่รายการ</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">รวมเป็นเงิน</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="day && filteredItems.length === 0">
            <td colspan="4" class="p-4 text-center">ไม่พบข้อมูลสำหรับวันที่: {{ day }}</td>
          </tr>
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
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import FetchApi from '@/fetch/FetchApi'
import { onMounted, ref, computed } from 'vue'

export default {
  setup() {
    const item = ref([])
    const day = ref(null)

    const Getmoney = async () => {
      const response = await FetchApi.GetMoney()
      item.value = response.data
      item.value.forEach((valve) => {
        const fullDateTime = valve.CreatedAt.split('T')
        const dateOnly = fullDateTime[0]
        valve.dateOnly = dateOnly
      })
      console.log(item.value)
    }

    const filteredItems = computed(() => {
      if (!day.value) return item.value
      return item.value.filter((valve) => valve.dateOnly === day.value)
    })

    onMounted(() => {
      Getmoney()
    })

    return {
      item,
      day,
      filteredItems,
    }
  },
}
</script>