import { defineStore } from 'pinia'

export const useUserOrder = defineStore('Userorder', {
  state: () => ({
    data: {
      customer_at_table: '',
      order: [],
    },
  }),
  actions: {
    addorder(customer_at_table, order) {
      this.data.customer_at_table = customer_at_table
      this.data.order.push(order)

      console.log({
        customer: this.data.customer_at_table,
        order:this.data.order
      })
    },
  },
})
