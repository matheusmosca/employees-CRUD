import Axios from 'axios'

const API =  Axios.create({
  baseURL: 'http://localhost:3030/api/v1'
  // baseURL: 'https://crudcrud.com/api/07623c902aa543bc80c9fb71173f0a39'
})

export default API 