<template>
  <div class="mx-36 my-72">
    <div class="grid grid-cols-3 gap-x-6 gap-y-24">

      <Card title="คิวอาร์โค๊ด" pathimage="/qr-code.png" />
      <Card title="ออเดอร์" pathimage="/menu.png"/>
      <Card title="เพิ่มเมนู" pathimage="/clipboard.png"/>
      <Card title="สรุปรายได้" pathimage="/budgeting.png"/>
      <Card title="จ่ายเงิน" pathimage="/payment.png"/>
    </div>
  </div>
</template>

<script>
import Card from '@/components/card.vue';

export default {
  name: 'App',
  components: {
    Card,
  },
};
</script>
