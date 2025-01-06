<template>
  <div class="flex justify-center items-center min-h-screen bg-gray-50">
    <div class="w-full max-w-md">
      <div class="bg-white rounded-lg shadow-xl p-8">
        <div class="text-center mb-8">
          <div class="inline-flex justify-center items-center w-16 h-16 bg-blue-100 rounded-full mb-4">

            <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8 text-blue-600" fill="none" viewBox="0 0 24 24"
              stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
            </svg>
          </div>
          <h1 class="text-2xl font-bold text-gray-900">เข้าสู่ระบบ</h1>
          <p class="text-gray-600 mt-2">กรุณากรอกข้อมูลเพื่อเข้าสู่ระบบ</p>
        </div>

        <form class="space-y-6" @submit="handleSubmit">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">
              ชื่อ
            </label>
            <input type="text" v-model="formState.username"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none transition-colors"
              placeholder="กรอกชื่อของคุณ" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">
              นามสกุล
            </label>
            <input type="text" v-model="formState.password"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none transition-colors"
              placeholder="กรอกนามสกุลของคุณ" />
          </div>

          <button type="submit"
            class="w-full bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 focus:ring-4 focus:ring-blue-200 transition-colors font-medium">
            เข้าสู่ระบบ
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import FetchApi from '@/fetch/FetchApi.vue';

import { reactive } from 'vue'
import { useRouter } from 'vue-router';

const router = useRouter();

const formState = reactive({
  username: '',
  password: ''
})

const handleSubmit = async (event) => {
  event.preventDefault();
  const username = formState.username
  const password = formState.password

  const url = 'http://localhost:8080/login';
  const payload = { username, password };
  try {
    const response = await FetchApi.methods.Login(url, payload);
    await localStorage.setItem('token', response.token);
    router.push('/MenuItems');
  } catch (error) {
    if (error.response && error.response.data) {
      alert(error.response.data.error || 'เกิดข้อผิดพลาดจากเซิร์ฟเวอร์');
    }
  }

}
</script>
