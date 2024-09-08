import { Route, RouterProvider, createBrowserRouter, createRoutesFromElements } from "react-router-dom";
import App from "../App";
import {Login, Register, Verify, Quessions} from '@pages'
import Layout from '@layout'


export default function Router(){
    const root = createBrowserRouter(
        createRoutesFromElements(
            <Route path="/" element={<App/>}>
                <Route index element={<Login/>} />
                <Route path="/register" element={<Register/>} />
                <Route path="/verify-password" element={<Verify/>} />
                <Route path="/dashboard/*" element={<Layout/>}>
                    <Route index element={<Quessions/>} />
                </Route>
            </Route>
        )
    )
    return <RouterProvider router={root} />
}