<template>
  <div class="grid grid-cols-2">
    <div class=" p-10">
      <form class="flex flex-col gap-y-8" @submit="onsubmitform">
        <div class="space-y-2">
          <label class="block text-gray-700 text-sm font-semibold">ชื่อเมนู</label>
          <input type="text" v-model="onsubmitdata.NameMenu"
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-200"
            placeholder="กรุณากรอกชื่อเมนู">

        </div>


        <div class="space-y-2">
          <label class="block text-gray-700 text-sm font-semibold">ราคา</label>
          <input type="text" v-model="onsubmitdata.Price"
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-200"
            placeholder="กรุณากรอกราคา เฉพาะตัวเลข">
        </div>


        <div class="space-y-2">
          <label class="block text-gray-700 text-sm font-semibold">สถานะ</label>
          <select v-model="onsubmitdata.Status"
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-200">
            <option value="">เลือกสถานะ</option>
            <option value="available">พร้อมขาย</option>
            <option value="unavailable">ไม่พร้อมขาย</option>
          </select>
        </div>


        <div class="space-y-2">
          <label class="block text-gray-700 text-sm font-semibold">หมวดหมู่</label>
          <input type="text" v-model="onsubmitdata.Category"
            class="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-200"
            placeholder="กรุณากรอกราคา เฉพาะตัวเลข">
        </div>

        <div class="space-y-2">
          <div class="flex flex-col justify-center items-center border border-collapse">
            <div class="p-10">
              <label htmlFor="file-upload"
                className="inline-flex flex-col items-center gap-2 cursor-pointer hover:text-blue-600 transition-colors duration-200">
                <!-- <Upload className="w-5 h-5" /> -->
                <span>อัพโหลดรูป</span>
                <h1>{{ fileuploadname }}</h1>
              </label>

              <input id="file-upload" type="file" accept="image/*" class="hidden" @change="onFileChange" />

            </div>
          </div>
        </div>


        <button type="submit"
          class="w-full bg-blue-500 text-white font-semibold py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition duration-200">
          บันทึกข้อมูล
        </button>
      </form>

    </div>
    <div>
      <table class="w-full border-collapse">
        <thead>
          <tr>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">รูปภาพ</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">ชื่อเมนู</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">ราคา</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">สถานะ</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">หมวดหมู่</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in getitem" :key="item.ID">
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">
              <img :src="`http://localhost:8080/Picture/${item.PictureMenu}`" class="w-20 h-20 object-cover rounded"
                :alt="item.NameMenu" />
            </th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">{{ item.NameMenu }}</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">{{ item.Price }}</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">{{ item.Status }}</th>
            <th scope="col" class="p-4 text-left border-b-2 border-gray-200">{{ item.Category }}</th>

          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import FetchApi from '@/fetch/FetchApi.vue';
import { onMounted, reactive, ref } from 'vue';

export default {
  setup() {
    const getitem = ref([]);
    const fileuploadname = ref('');
    const onsubmitdata = reactive({
      file: "",
      NameMenu: "",
      Price: "",
      Status: "",
      Category: "",
    })

    const onsubmitform = async (event) => {
      event.preventDefault();
      const formdata = await new FormData()
      formdata.append("file",onsubmitdata.file)
      formdata.append("namemenu",onsubmitdata.NameMenu)
      formdata.append("price",onsubmitdata.Price)
      formdata.append("status",onsubmitdata.Status)
      formdata.append("category",onsubmitdata.Category)
      console.log(formdata)
      const res = await FetchApi.methods.addmenu(formdata)
      console.log(res)
      getdata()
    }

    const getdata = async () => {
      const res = await FetchApi.methods.getmenu();
      getitem.value = res.data.data;
    };



    const onFileChange = (event) => {
      const file = event.target.files[0];

      if (file) {
        fileuploadname.value = file.name;
        onsubmitdata.file = file;
        console.log(onsubmitdata.file)
      }
    };

    onMounted(() => {
      getdata();
    });

    return {
      getitem,
      getdata,
      fileuploadname,
      onFileChange,
      onsubmitdata,
      onsubmitform
    };
  }
}
</script>
