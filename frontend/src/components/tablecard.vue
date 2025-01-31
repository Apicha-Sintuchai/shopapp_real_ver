    <template>
      <div class="mx-20 my-8">
        <div class="flex justify-end items-center">
          <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" @click="toggleAddDataAlert">
            เพิ่มโต๊ะ
          </button>
        </div>
      </div>
      <div class="mx-36 my-48">
        <div class="relative grid grid-cols-4 gap-8">
          <div v-for="item in title" :key="item.id">
            <div class="max-w-72 bg-gray-300 rounded-xl p-16 shadow-lg">
              <div class="flex flex-col justify-center items-center gap-y-12">
                <img src="/qr-code.png" alt="QR Code" class="w-24 h-24" />
                <h1 class="font-bold text-3xl">{{ item.table_number }}</h1>
                <div class="flex justify-between items-center  ">
                  <button @click="printQRCode(item.table_number)" class="bg-blue-500 text-white px-4 py-2 rounded-l-xl w-16">
                    ปริ้น
                  </button>
                  <button @click="() => toggleUpdateAlert(item)" class="bg-yellow-500 text-white px-4 py-2  w-24">
                    แก้ไข
                  </button>
                  <button @click="deleteById(item.id)" class="bg-red-500 text-white px-4 py-2 rounded-r-xl w-16 ">
                    ลบ
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>

    <script>
    import { ref, reactive, onMounted } from "vue";
    import QRCode from "qrcode";
    import FetchApi from "@/fetch/FetchApi";
    import Swal from "sweetalert2";
    export default {
      name: "TableCard",
      setup() {
        const title = ref([]);
        const formdata = reactive({
          table_number: "",
        });

        const fetchTables = async () => {
         
            const res = await FetchApi.GetTable();
            console.log(res)
            title.value =  res.data.data;
      
        };

        const printQRCode = async (tableNumber) => {
          try {
            const datauql = "http://localhost/customer/" +tableNumber
            const qrCodeDataUrl = await QRCode.toDataURL(datauql);
            const link = document.createElement("a");
            link.href = qrCodeDataUrl;
            link.download = `QRCode-${tableNumber}.png`;
            link.click();
          } catch (err) {
            console.error("Error generating QR code:", err);
          }
        };

        const toggleAddDataAlert = () => {
          Swal.fire({
            title: "เพิ่มโต๊ะ",
            html: '<input id="table-input" class="border-4 w-full" placeholder="ใส่ชื่อโต๊ะ" />',
            showCancelButton: true,
            confirmButtonText: "ยืนยัน",
            cancelButtonText: "ยกเลิก",
            preConfirm: () => {
              const inputValue = document.getElementById("table-input").value;
              if (!inputValue) {
                Swal.showValidationMessage("กรุณากรอกชื่อโต๊ะ");
              }
              return inputValue;
            },
          }).then(async (result) => {
            if (result.isConfirmed) {
              formdata.table_number = result.value;
              const res = await FetchApi.AdddataTable(formdata.table_number)
              console.log(res)
              fetchTables()

            }
          });
        };

        const toggleUpdateAlert = (item) => {
          Swal.fire({
            title: "แก้ไขข้อมูล",
            html: `<input id="table-input-update" class="border-4 w-full" placeholder="${item.table_number}" />`,
            showCancelButton: true,
            confirmButtonText: "ยืนยัน",
            cancelButtonText: "ยกเลิก",
            preConfirm: () => {
              const inputValue = document.getElementById("table-input-update").value;
              if (!inputValue) {
                Swal.showValidationMessage("กรุณากรอกชื่อโต๊ะ");
              }
              return inputValue;
            },
          }).then(async (result) => {
            if (result.isConfirmed) {
              formdata.table_number = result.value;
              console.log("แก้ไขข้อมูล:", formdata.table_number);
              const res = await FetchApi.UpdateTable(item.ID,formdata.table_number)
              console.log(res)
              fetchTables()
            }
          });
        };


        const deleteById = async (id) => {
          try {
            const res = await FetchApi.DeleteTable(id);
            console.log("Deleted table:", res);
            fetchTables();
          } catch (error) {
            console.error("Error deleting table:", error);
          }
        };

        onMounted(() => {
          
          setTimeout(fetchTables, 2000);
         
        });

        return {
          title,
          formdata,
          printQRCode,
          toggleAddDataAlert,
          toggleUpdateAlert,
          deleteById,
        };
      },
    };
    </script>
