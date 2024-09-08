import React, { useEffect, useState } from 'react';
import { Flex, Input, Typography } from 'antd';
import { ToastContainer, toast } from 'react-toastify';
import './style.scss';
import http from '../../config';
import { useNavigate } from 'react-router-dom';

type OTPProps = React.ComponentProps<typeof Input.OTP>;

const { Title } = Typography;

const App: React.FC = () => {
  const [sekund, setSecund] = useState(59);
  const [minut, setMinut] = useState(4);
  const navigate = useNavigate();

  const onChange: OTPProps['onChange'] = async (text: any) => {
    if (text.length === 6) {
      try {
        const response = await http.post('/confirm-registration', { code: text, email: localStorage.getItem('email') });
        if (response?.status === 200) {
          toast.success('Kod muvaffaqiyatli tasdiqlandi !', { autoClose: 1200 });
          setTimeout(() => {
            navigate('/');
          }, 1500);
        }
        console.log(response);
      } catch (err) {
        toast.error('Kodda xatolik mavjud qayta tekshirib ko\'ring !', { autoClose: 1200 });
        console.log(err);
      }
    }
  };

  async function resetVerify(){
    setMinut(4)
    setSecund(59)
    const response = await http.post('/resverify',  {email: localStorage.getItem('email')})
    if(response.status == 200){
      toast.success('Sizga qayta Cod yuborildi!', { autoClose: 1200 });
    }
  }

  useEffect(() => {
    const handleBeforeUnload = (event: BeforeUnloadEvent) => {
      event.preventDefault()
      navigate('/'); 
    };
    window.addEventListener('beforeunload', handleBeforeUnload);
  }, [navigate]);

  useEffect(() => {
    const intervalId = setInterval(() => {
      if (sekund === 1) {
        if (minut === 0) {
          clearInterval(intervalId);
        } else {
          setMinut((prevMinut) => prevMinut - 1);
          setSecund(59);
        }
      } else {
        setSecund((prevSecund) => prevSecund - 1);
      }
    }, 1000);

    return () => clearInterval(intervalId);
  }, [sekund, minut]);

  return (
    <div>
      <ToastContainer />
      <div className="verify" style={{ width: '100%', maxWidth: '400px', padding: 30, margin: '0 auto', marginTop: 200 }}>
        <Flex gap="middle" align="flex-start" vertical>
          <Title level={5}>Emailingizga tasdiqlash kodini yubordik. <br /> Kodni kiriting!</Title>
          <Input.OTP onChange={(text) => onChange(text)} length={6} />
        </Flex>
        <p  onClick={() => resetVerify()} style={{textAlign: 'center', marginTop: 20, fontWeight: 600, color: 'blue', cursor: 'pointer'}}>{minut == 0 && sekund == 1 ? 'Qaytadan yuborish' : `${minut}:${sekund.toString().padStart(2, '0')}`}</p>
      </div>
    </div>
  );
};

export default App;
