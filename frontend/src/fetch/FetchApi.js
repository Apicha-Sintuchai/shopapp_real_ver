import axios from 'axios';

const FetchApi = {
  async getToken() {
    return sessionStorage.getItem("token");
  },

  async Login(url, payload) {
    const response = await axios.post(url, payload);
    return response.data;
  },

  async GetTable() {
    const token = await this.getToken();
    const response = await axios.get("http://localhost:8080/admin/tables", {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response;
  },

  async AdddataTable(payload) {
    const token = await this.getToken();
    const response = await axios.post(
      "http://localhost:8080/admin/tables",
      { table_number: payload },
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response;
  },

  async UpdateTable(id, payload) {
    const token = await this.getToken();
    const response = await axios.put(
      `http://localhost:8080/admin/tables/whereidupdate/${id}`,
      { table_number: payload },
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response;
  },

  async GetOrder() {
    const token = await this.getToken();
    const response = await axios.get("http://localhost:8080/admin/customer_orders", {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response;
  },

  async donemenu(id) {
    const token = await this.getToken();
    const response = await axios.put(
      `http://localhost:8080/admin/customer_orders/status/${id}`,
      {},
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response;
  },

  async getmenu() {
    const token = await this.getToken();
    const response = await axios.get("http://localhost:8080/admin/menus", {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response;
  },

  async addmenu(data) {
    const token = await this.getToken();
    const response = await axios.post("http://localhost:8080/admin/menus", data, {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response;
  },

  async GetMoney() {
    const token = await this.getToken();
    const response = await axios.get("http://localhost:8080/admin/customer_orders/GetMoney", {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response;
  },

  async Getdata_unpayment() {
    const token = await this.getToken();
    const response = await axios.get(
      "http://localhost:8080/admin/customer_orders/Getdata_unpayment",
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response;
  },

  async updatepayment(id) {
    const token = await this.getToken();
    const response = await axios.put(
      `http://localhost:8080/admin/customer_orders/updatepayment/${id}`,
      {},
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response;
  },

  async customer_getmenu() {
    const response = await axios.get("http://localhost:8080/customer/menus");
    return response;
  },
};

export default FetchApi;
