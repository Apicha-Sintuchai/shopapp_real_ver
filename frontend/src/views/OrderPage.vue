<template>
  <div class="mx-36 my-48">
    <div class="relative grid grid-cols-4 gap-8">
      <div v-for="item in title" :key="item.ID">
        <div class="max-w-72 bg-gray-300 rounded-xl p-4 shadow-lg">
          <div class="flex flex-col justify-center items-center gap-y-12">
            <img src="/menu.png" alt="QR Code" class="w-24 h-24" />
            <h1 class="font-bold text-3xl">{{ item.customer_at_table }}</h1>
            <h1>สั่ง ณ เวลา {{ item.dateOnly }} เวลา {{ item.timeOnly }}</h1> <!-- แสดงวันที่และเวลา -->
            <div class="flex flex-col justify-center items-center gap-3">
              <button class="bg-blue-500 text-white px-4 py-2 rounded" @click="openModal(item)">
                ดูรายการ
              </button>
              <button class="bg-red-500 text-white px-4 py-2 rounded" @click="changedonestatus(item.ID)">
                เสร็จรายการ
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal -->
    <div v-if="isModalOpen" class="fixed inset-0 bg-gray-800 bg-opacity-50 flex justify-center items-center">
      <div class="bg-white rounded-xl p-6 w-96">
        <h2 class="text-xl font-bold mb-4">รายละเอียดรายการ</h2>

        <p v-if="selectedItem">โต๊ะที่ &nbsp;{{ selectedItem.customer_at_table }}</p>
        <p class="font-bold text-xs">รายการ</p>
        <div v-for="orderitem in selectedItem.orders" :key="orderitem.ID">
          <div class="flex  gap-8">
            <p>{{ orderitem.name_menu }}</p>
          
          <p v-if="orderitem.option">เพิ่มเติม</p>
          <p v-if="`${orderitem.option != null} `">{{ orderitem.option }}</p>
          </div>
          
          
        </div>

        <button class="bg-red-500 text-white px-4 py-2 rounded mt-4" @click="closeModal">
          ปิด
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import FetchApi from '@/fetch/FetchApi.vue';
import { onMounted, ref } from 'vue';

export default {
  setup() {
    const title = ref([]);

    const isModalOpen = ref(false);
    const selectedItem = ref(null);

    const GetData = async () => {
      try {
        const res = await FetchApi.methods.GetOrder();
        title.value = res.data;
        title.value.forEach((valve) => {
          const fullDateTime = valve.CreatedAt.split("T");
          const dateOnly = fullDateTime[0];
          const fullTime = fullDateTime[1];
          const timeOnly = fullTime.split(":").slice(0, 2).join(":");
          valve.dateOnly = dateOnly;
          valve.timeOnly = timeOnly;
        });

      } catch (error) {
        console.log(error);
      }
    };

    const openModal = (item) => {
      selectedItem.value = item;
      isModalOpen.value = true;
    };

    const closeModal = () => {
      isModalOpen.value = false;
      selectedItem.value = null;
    };

    const changedonestatus = async (id) => {
      await FetchApi.methods.donemenu(id)
      GetData()
    }

    onMounted(() => {
      GetData();
    });

    return {
      title,
      isModalOpen,
      selectedItem,
      GetData,
      openModal,
      closeModal,
      changedonestatus,
    };
  },
};
</script>
