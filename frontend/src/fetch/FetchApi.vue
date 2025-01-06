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
  },
};
</script>
