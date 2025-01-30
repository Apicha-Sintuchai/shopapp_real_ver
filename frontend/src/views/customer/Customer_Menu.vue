<template>
  <div>
    <div class="flex flex-col gap-20 p-4"
    v-for="items in getitem"
    :key="items.ID"
    >
      <!-- header -->
      <div
        class="flex gap-8 border border-1 border-solid overflow-x-auto no-scrollbar sm:overflow-hidden"
      >
        <h1>{{ items.Category }}</h1>
      </div>

       <!-- body -->
       <div>
          <h1>{{ items.NameMenu }}</h1>
        </div>
    </div>
  </div>
</template>
<script>
import FetchApi from '@/fetch/FetchApi.vue'
import { onMounted, ref } from 'vue'

export default {
  setup() {
    const getitem = ref([])

    const getdata = async () => {
      const res = await FetchApi.methods.customer_getmenu()
       getitem.value =  await res.data.data

      console.log(getitem.value)
    }

    onMounted(() => {
      getdata()
      getitem
    })

    return {
      getitem,
    }
  },
}
</script>
