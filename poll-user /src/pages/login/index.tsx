import { toast, ToastContainer } from "react-toastify";
import { Button, Form, Input } from "antd";
import FormItem from "antd/es/form/FormItem";
import { LockOutlined, MailOutlined } from "@ant-design/icons";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import "./style.scss";
import http from "../../config";
import { setCookies } from "../../utils/cocies";

function Login() {
  const [load, setLoadi] = useState(false)
  const navigate = useNavigate()


  async function handleSubmit(values: any) {
    try{
      const response = await http.post('/login', values)
      if(response?.status == 200){
        setCookies('access_token', response?.data?.access_token)
        toast.success('Tizimga kirishga ruhsat berildi !', {autoClose: 1300})
        setTimeout(() => {
          navigate('/dashboard')
        }, 2000);
      }
    }catch(err:any){
      toast.error(err?.response?.data?.error == 'email' ? 'Email xato kiritildi' : 'Password xato kiritild')
    }
    setLoadi(true)
    setTimeout(() => {
      setLoadi(false)
    }, 3000);
  }

  return (
    <>
      <ToastContainer />

      <div className="login">
        <div className="login-form">
          <div className="login-title">
            {/* <img
              src="https://cdn.iconscout.com/icon/free/png-256/free-leetcode-3521542-2944960.png?f=webp"
              alt=""
            /> */}
            <h1>Psixolog</h1>
          </div>

          <div className="login-body">
            <Form onFinish={(values) => handleSubmit(values)}>
              <p>Email</p>
              <FormItem
                name="email"
                hasFeedback
                rules={[
                  {
                    type: "email",
                    required: true,
                    message: "Iltimos Emailni to'g'ri formatda kiriting!",
                  },
                ]}
                
              >
                <Input prefix={<MailOutlined/>} size="large" placeholder="Email" name="email" />
              </FormItem>
              
              <div style={{display: 'flex', alignItems: 'center', justifyContent: 'space-between', fontWeight: 600}}>
                <p>Password</p>
              </div>
              <FormItem
                name="password"
                hasFeedback
                rules={[
                  {
                    required: true,
                    message: "Iltimos parolingizni kiriting !",
                  },
                ]}
              >
                <Input.Password prefix={<LockOutlined/>} size="large" name="password" placeholder="password" />
              </FormItem>

              <FormItem>
                <Button loading={load} htmlType="submit" type="primary">
                  Submit
                </Button>
              </FormItem>

              <div className="login-tooltip"></div>

              <Link className="login-link" to={'/register'}>
                  Register
              </Link>
            </Form>
          </div>
        </div>
      </div>
      
    </>
  );
}

export default Login;