'use client'
import { Button } from '@/components/ui/button'
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import React from 'react'
import { SubmitHandler, useForm } from 'react-hook-form'
import { Login } from '../(action)/fetch'

type FormValues = {
  username: string
  password: string
}

const Page = () => {
  const { register, handleSubmit, formState: { errors } } = useForm<FormValues>()

  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    const res = await Login(data.username, data.password)
    
    if (res.success) {
      localStorage.setItem('token', res.data.token)
    }
    else {
      console.log("login failed")
    }


  }

  return (
    <div className='flex justify-center items-center h-screen'>
      <Card className="w-[350px]">
        <CardHeader>
          <CardTitle>Login</CardTitle>
          <CardDescription>Use your username and password</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit(onSubmit)}>
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="username">Username</Label>
                <Input 
                  id="username" 
                  placeholder="Enter your username" 
                  {...register('username', { required: 'Username is required' })} 
                />
                {errors.username && <p className="text-red-500 text-sm">{errors.username.message}</p>}
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="password">Password</Label>
                <Input 
                  id="password" 
                  type="password" 
                  placeholder="Enter your password" 
                  {...register('password', { required: 'Password is required' })} 
                />
                {errors.password && <p className="text-red-500 text-sm">{errors.password.message}</p>}
              </div>
            </div>
            <CardFooter className="flex justify-between mt-4">
              <Button type="submit">Login</Button>
            </CardFooter>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}

export default Page
