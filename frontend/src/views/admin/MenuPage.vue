<template>
  <div class="min-h-screen">
    <div class="grid grid-cols-1 lg:grid-cols-2 h-full">
      <!-- Form Section -->
      <div class="p-6 lg:p-10">
        <form class="flex flex-col gap-y-8" @submit="onsubmitform">
          <div class="space-y-2">
            <label class="block text-gray-950 text-sm font-semibold">ชื่อเมนู</label>
            <input
              type="text"
              v-model="onsubmitdata.NameMenu"
              class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-200"
              :class="{ 'border-red-500': v$.NameMenu.$error }"
              placeholder="กรุณากรอกชื่อเมนู"
            />
            <div class="text-red-500 text-sm" v-if="v$.NameMenu.$error">กรุณากรอกชื่อเมนู</div>
          </div>

          <div class="space-y-2">
            <label class="block text-gray-700 text-sm font-semibold">ราคา</label>
            <input
              type="text"
              v-model="onsubmitdata.Price"
              class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-200"
              :class="{ 'border-red-500': v$.Price.$error }"
              placeholder="กรุณากรอกราคา เฉพาะตัวเลข"
            />
            <div class="text-red-500 text-sm" v-if="v$.Price.$error">กรุณากรอกราคา</div>
          </div>

          <div class="space-y-2">
            <label class="block text-gray-700 text-sm font-semibold">สถานะ</label>
            <select
              v-model="onsubmitdata.Status"
              class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-200"
              :class="{ 'border-red-500': v$.Status.$error }"
            >
              <option value="">เลือกสถานะ</option>
              <option value="available">พร้อมขาย</option>
              <option value="unavailable">ไม่พร้อมขาย</option>
            </select>
            <div class="text-red-500 text-sm" v-if="v$.Status.$error">กรุณาเลือกสถานะ</div>
          </div>

          <div class="space-y-2">
            <label class="block text-gray-950 text-sm font-semibold">หมวดหมู่</label>
            <input
              type="text"
              v-model="onsubmitdata.Category"
              class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-200"
              :class="{ 'border-red-500': v$.Category.$error }"
              placeholder="กรุณากรอกหมวดหมู่"
            />
            <div class="text-red-500 text-sm" v-if="v$.Category.$error">กรุณากรอกหมวดหมู่</div>
          </div>

          <div class="space-y-2">
            <div
              class="flex flex-col justify-center items-center border border-collapse"
              :class="{ 'border-red-500': v$.file.$error }"
            >
              <div class="p-10">
                <label
                  htmlFor="file-upload"
                  class="inline-flex flex-col items-center gap-2 cursor-pointer hover:text-blue-600 transition-colors duration-200"
                >
                  <span>อัพโหลดรูป</span>
                  <h1>{{ fileuploadname }}</h1>
                </label>
                <input
                  id="file-upload"
                  type="file"
                  accept="image/*"
                  class="hidden"
                  @change="onFileChange"
                />
              </div>
            </div>
            <div class="text-red-500 text-sm" v-if="v$.file.$error">กรุณาอัพโหลดรูปภาพ</div>
          </div>

          <button
            type="submit"
            class="w-full bg-blue-500 text-white font-semibold py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition duration-200"
          >
            บันทึกข้อมูล
          </button>
        </form>
      </div>

      <!-- Table Section -->
      <div class="bg-slate-100 min-h-screen overflow-hidden">
        <div class="overflow-x-auto">
          <div class="max-h-screen overflow-y-auto">
            <table class="w-full border-collapse">
              <thead class="sticky top-0 bg-slate-100">
                <tr>
                  <th scope="col" class="p-4 text-left border-b-2 border-gray-200">รูปภาพ</th>
                  <th scope="col" class="p-4 text-left border-b-2 border-gray-200">ชื่อเมนู</th>
                  <th scope="col" class="p-4 text-left border-b-2 border-gray-200">ราคา</th>
                  <th scope="col" class="p-4 text-left border-b-2 border-gray-200">สถานะ</th>
                  <th scope="col" class="p-4 text-left border-b-2 border-gray-200">หมวดหมู่</th>
                  <th scope="col" class="p-4 text-left border-b-2 border-gray-200">action</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in getitem" :key="item.ID">
                  <td scope="col" class="p-4 text-left border-b-2 border-gray-200">
                    <img
                      :src="`http://localhost:8080/Picture/${item.PictureMenu}`"
                      class="w-20 h-20 object-cover rounded"
                      :alt="item.NameMenu"
                    />
                  </td>
                  <td scope="col" class="p-4 text-left border-b-2 border-gray-200">
                    {{ item.NameMenu }}
                  </td>
                  <td scope="col" class="p-4 text-left border-b-2 border-gray-200">
                    {{ item.Price }}
                  </td>
                  <td scope="col" class="p-4 text-left border-b-2 border-gray-200">
                    {{ item.Status }}
                  </td>
                  <td scope="col" class="p-4 text-left border-b-2 border-gray-200">
                    {{ item.Category }}
                  </td>
                  <td scope="col" class="p-4 text-left border-b-2 border-gray-200 rounded">
                    <button
                    @click="detelemenu(item.ID)"
                      class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded-full"
                    >
                      ลบ
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import FetchApi from '@/fetch/FetchApi'
import { onMounted, reactive, ref } from 'vue'
import { useVuelidate } from '@vuelidate/core'
import { required, numeric } from '@vuelidate/validators'

const getitem = ref([])
const fileuploadname = ref('')
const onsubmitdata = reactive({
  file: '',
  NameMenu: '',
  Price: '',
  Status: '',
  Category: '',
})

const rules = {
  file: { required },
  NameMenu: { required },
  Price: { required, numeric },
  Status: { required },
  Category: { required },
}

const v$ = useVuelidate(rules, onsubmitdata)

const onsubmitform = async (event) => {
  event.preventDefault()
  const isValid = await v$.value.$validate()

  if (!isValid) {
    console.log('Validation errors:', v$.value.$errors)
    return
  }

  try {
    const formData = new FormData()
    formData.append('file', onsubmitdata.file)
    formData.append('namemenu', onsubmitdata.NameMenu)
    formData.append('price', onsubmitdata.Price)
    formData.append('status', onsubmitdata.Status)
    formData.append('category', onsubmitdata.Category)

    const response = await FetchApi.addmenu(formData)
    console.log('Form submitted successfully:', response)

    v$.value.$reset()
    Object.keys(onsubmitdata).forEach((key) => (onsubmitdata[key] = ''))
    fileuploadname.value = ''

    await getdata()
  } catch (error) {
    console.error('Error submitting form:', error)
  }
}

const getdata = async () => {
  try {
    const res = await FetchApi.getmenu()
    getitem.value = res.data.data
  } catch (error) {
    console.error('Error fetching menu data:', error)
  }
}

const onFileChange = (event) => {
  const file = event.target.files[0]
  if (file) {
    fileuploadname.value = file.name
    onsubmitdata.file = file
  }
}


const detelemenu = async(id) => {
  try {
    const res = await FetchApi.DeleteMenu(id)
    console.log(res)
    getdata()
  } catch (error) {
    console.log(error)
  }
  
}
onMounted(() => {
  getdata()
})
</script>

