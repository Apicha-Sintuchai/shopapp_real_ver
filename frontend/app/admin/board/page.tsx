import { Card, CardHeader } from '@/components/ui/card';
import Link from 'next/link';

import React from 'react';
import { FaQrcode } from "react-icons/fa";
import { GiFoodTruck } from "react-icons/gi";
import { MdMenuBook } from "react-icons/md";
import { TbReport } from "react-icons/tb";

const page = () => {
    const itemdata = [
        {
            Link: "/admin/qrcode",
            Name: "คิวอาร์โค๊ด",
            icon: <FaQrcode size={55} />,
        },
        {
            Link: "/admin/order",
            Name: "ออเดอร์",
            icon: <GiFoodTruck size={55} />,
        },
        {
            Link: "/admin/menu",
            Name: "เพิ่มเมนู",
            icon: <MdMenuBook size={55} />,
        },
        {
            Link: "/admin/report",
            Name: "report",
            icon: <TbReport size={55} />,
        },
    ];

    return (
        <>
            <div className="mx-48 my-60">
                <div className="grid grid-cols-3 gap-16">
                    {itemdata.map((item, index) => (
                        <Card key={index}>
                            <CardHeader>
                                <Link href={item.Link}>
                                    <div className="flex justify-between items-center">
                                        {item.Name}
                                        {item.icon}
                                    </div>
                                </Link>
                            </CardHeader>
                        </Card>
                    ))}
                </div>
            </div>
        </>
    );
};

export default page;
