<template>
  <div class="mx-24 my-72">
    <div class="grid grid-cols-1 gap-x-6 gap-y-24 sm:grid-cols-3">
      <RouterLink to="/Qrcodepage">
        <Card title="คิวอาร์โค๊ด" pathimage="/qr-code.png" />
      </RouterLink>
      <RouterLink to="/Orderpage">
        <Card title="ออเดอร์" pathimage="/menu.png" />
      </RouterLink>
      <RouterLink to="/MenuPage">
        <Card title="เพิ่มเมนู" pathimage="/clipboard.png" />
      </RouterLink>
      <RouterLink to="/Recivepage">
        <Card title="สรุปรายได้" pathimage="/budgeting.png" />
      </RouterLink>
      <RouterLink to="/Paymentpage">
        <Card title="จ่ายเงิน" pathimage="/payment.png" />
      </RouterLink>
    </div>
  </div>
</template>

<script>
import Card from '@/components/card.vue'

export default {
  name: 'App',
  components: {
    Card,
  },
}
</script>
