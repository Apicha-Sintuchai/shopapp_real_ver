'use server'
import axios from 'axios';




export const GetQrcode = async () => {
  
        const res = await fetch("http://localhost:8080/admin/tables",{
            credentials: 'include',
        });
        
        if (!res.ok) {
            return {
                success: false,
                error: `HTTP Error: ${res.status} ${res.statusText}`,
            };
        }
        const data = await res.json();
        return {
            success: true,
            data,
        };

};




export const Login = async (username: string, password: string) => {
    interface LoginResponse {
        token: string;
        
    }

    const res = await axios.post<LoginResponse>(
        'http://localhost:8080/login', 
        { username, password },
        {
            headers: {
                'Content-Type': 'application/json; charset=utf-8',
            },
            withCredentials: true, 
        }
    );

    
 

    return {
        success: true,
        data: res.data,
    };
};