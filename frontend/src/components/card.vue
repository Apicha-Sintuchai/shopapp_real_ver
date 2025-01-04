<template>
  <div class="bg-gray-100 p-4 rounded-lg shadow-lg">
    <div class="flex justify-between items-center">
      <h1>{{ title }}</h1>
      <img :src="pathimage" alt="icon"  class="w-12 h-12"/>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CardComponent',
  props: {
    title: {
      type: String,
      default: '',
    },
    pathimage: {
      type: String,
      default: '',
    },
  },
};
</script>
