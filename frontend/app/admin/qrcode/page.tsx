'use client'
import React, { useEffect } from 'react'
import { Button } from "@/components/ui/button"
import { GetQrcode } from '@/app/(action)/fetch'

type Props = {}

const page = (props: Props) => {

  const fetchQrcodeData = async () => {
    const result = await GetQrcode();

    if (result.success) {
        console.log("Data:", result.data);
    } else {
        console.error("Error:", result.error);
    }
};

useEffect(() => {
    fetchQrcodeData();
}
, []);

  return (
    <div className='mx-10 my-10'>
      <div className='flex justify-between items-center '>
        <div>
          <h1 className='text-3xl font-bold'>คิวอาร์โค๊ด</h1>
        </div>
        <div>
        <Button variant="outline" className='shadow-lg rounded-full '>Button</Button>
        </div>
      </div>
    </div>

  )
}

export default page