  <template>
    <div class="p-6">
      <div class="overflow-x-auto no-scrollbar p-4 rounded-xl shadow-lg">
        <div class="flex w-max gap-4 sm:gap-6">
          <span
            v-for="category in categories"
            :key="category"
            class="px-6 py-2 bg-blue-500 text-white rounded-full shadow-md cursor-pointer transition-transform hover:scale-105"
          >
            {{ category }}
          </span>
        </div>
      </div>
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 mt-8">
        <div
          v-for="item in getitem"
          :key="item.ID"
          class="p-4 bg-white rounded-xl shadow-lg transform transition-transform hover:scale-105"
        >
          <img
            :src="`http://localhost:8080/Picture/${item.PictureMenu}`"
            class="w-32 h-32 object-cover rounded-lg mx-auto"
            :alt="item.NameMenu"
          />
          <div class="text-center mt-3">
            <h2 class="text-lg font-semibold">{{ item.NameMenu }}</h2>
            <p class="text-gray-600">{{ item.Category }}</p>
            <p class="text-green-500 font-bold text-lg mt-2">฿{{ item.Price }}</p>
            <button @click="orderincome(item)">เพิ่มเมนู</button>
          </div>
        </div>
      </div>

      <div v-if="(orderitem > 0)" class="p-4 shadow-xl mt-4 rounded-xl bg-purple-400">
        <div class="flex justify-between items-center relative">
          <h1>เมนูที่ท่านเลือก</h1>

          <div class="relative w-10 h-10">
            <router-link to="/order_cart" >
            <img src="/shopping-cart.png" alt="shopping-cart.png" class="w-8 h-8" />
            <span
              class="absolute -top-2 -right-1 bg-red-500 text-white text-xs rounded-full w-5 h-5 flex items-center justify-center shadow-lg"
              >{{ orderitem }}</span
            >
          </router-link>
          </div>
        </div>
      </div>
    </div>
  </template>

  <script>
  import FetchApi from '@/fetch/FetchApi'
  import { useUserOrder } from '@/stores/counter'
  import Swal from 'sweetalert2'
  import { computed, onMounted, ref } from 'vue'
  import { useRoute } from 'vue-router'

  export default {
    setup() {
      const router = useRoute()
      const getitem = ref([])
      const categories = ref([])
      const userOrder = useUserOrder()
      const customer_at_table = router.params.id
      const orderitem = computed(() => userOrder.getorderlenght())
      const getdata = async () => {
        const res = await FetchApi.customer_getmenu()
        getitem.value = res.data.data
        categories.value = [...new Set(getitem.value.map((item) => item.Category))]
      }
      
      const orderincome = async (orderitem) => {
      
        const { value } = await Swal.fire({
          title: 'ต้องการเพิ่มอะไร ถ้าไม่สามารถกดยืนยันได้เลย',
          input: 'text',
          inputPlaceholder: '',
        })

        const packagedata = {
          image : orderitem.PictureMenu,
          name_menu: orderitem.NameMenu,
          price: orderitem.Price,
          option: value,
        }

        userOrder.addorder(customer_at_table, packagedata)
      }

      onMounted(() => {
        getdata()
      })

      return {
        getitem,
        categories,
        userOrder,
        orderincome,
        orderitem
      }
    },
  }
  </script>
