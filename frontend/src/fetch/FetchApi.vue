<script>
import axios from 'axios';
const token = localStorage.getItem('token');
export default {
  name: 'FetchApi',
  methods: {
    async Login(url, payload) {
      const response = await axios.post(url, payload);
      return response.data;
    },
    async GetTable() {
      const response = await axios.get('http://localhost:8080/admin/tables', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      return response;
    },
    async AdddataTable(payload) {
      const response = await axios.post(`http://localhost:8080/admin/tables`, {
        "table_number": payload
      }, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })
      return response
    },
    async UpdateTable(id, payload) {
      const response = await axios.put(`http://localhost:8080/admin/tables/whereidupdate/${id}`, {
        "table_number": payload
      }, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })
      return response


    },
    async GetOrder() {
      const response = await axios.get("http://localhost:8080/admin/customer_orders",
        {
          headers: {
            Authorization: `Bearer ${token}`
          }
        }
      )
      return response
    },

    async donemenu(id) {
      const response = await axios.put(`http://localhost:8080/admin/customer_orders/status/${id}`, {}, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })
      return response
    }
  },
};
</script>
