import axios from "axios";
import { defineStore } from "pinia";

export const useRecordStore = defineStore({
  id: "record",
  state: () => ({
    records: [],
    activeRecords: [],
    activeRecord: {},
    statusMap: [
      { val: 1, label: "提议" },
      { val: 2, label: "通过" },
      { val: 3, label: "完成" },
      { val: 4, label: "已弃用" },
      { val: 5, label: "已取代" },
    ],
    loading: false,
  }),
  getters: {
    getStatusLabelById: state => {
      return (id: number) => state.statusMap.find(item => item.val === id)?.label
    },
  },
  actions: {
    async loadRecordByID(id: any) {
      this.loading = true;
      const res = await axios.get(`/api/record/${id}`);
      this.activeRecords = res.data?.data;
      this.loading = false;
      return res;
    },
    async createOrUpdateRecord(data: any) {
      this.loading = true;
      const res = await axios.post("/api/record", data);
      this.loading = false;
      return res;
    },
    async loadDetail(id: any) {
      this.loading = true;
      const res = await axios.get(`/api/records/${id}`);
      this.activeRecord = res.data?.data
      this.loading = false;
      return res;
    },
    async updateStatus(id: any, status: any) {
      this.loading = true;
      // TODO: new api
      const { data } = await this.loadDetail(id);
      const res = await this.createOrUpdateRecord({ ...data.data, status });
      this.loading = false;
      return res;
    },
  },
});
