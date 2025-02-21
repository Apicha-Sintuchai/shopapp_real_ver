import axios from 'axios'
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
export const useUserOrder = defineStore('Userorder', {
  state: () => ({
    data: {
      customer_at_table: '',
      order: [],
    },
  }),
  getters: {
    orderDetails: (state) => ({
      customer_at_table: state.data.customer_at_table,
      order: state.data.order,
    }),
  },
  actions: {
    addorder(customer_at_table, order) {
      this.data.customer_at_table = customer_at_table
      this.data.order.push(order)
    },
    removeOrderItem(itemName) {
      this.data.order = this.data.order.filter((item) => item.name_menu !== itemName)
    },

    getorderlenght() {
      return this.data.order.length
    },

    getorder() {
      return {
        customer_at_table: this.data.customer_at_table,
        order: this.data.order,
      }
    },

    async sendorder() {
      const res = await axios.post('http://localhost:8080/customer/customer_orders', {
        customer_at_table: this.data.customer_at_table,
        orders: this.data.order,
      })

      console.log(res.data)
    },
  },
})
