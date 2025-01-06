import axios from 'axios'

const API_URL = 'http://localhost:8080/admin'
const tokn = localStorage.getItem('token')
export const fetchTables = async () => {
  return await axios.get(`${API_URL}/tables`,{
    headers: {
      'Authorization': `Bearer ${tokn}`
    }
  })
}
