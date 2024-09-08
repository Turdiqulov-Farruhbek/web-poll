import { Route, RouterProvider, createBrowserRouter, createRoutesFromElements } from "react-router-dom";
import App from "../App";
import {Login, Home, Quession } from '@pages'
import Layout from '@layout'

export default function Router(){
    const root = createBrowserRouter(
        createRoutesFromElements(
            <Route path="/" element={<App/>}>
                <Route index element={<Login/>} />
                <Route path="/home" element={<Home/>} />
                <Route path='/dashboard/*' element={<Layout/>}>
                    <Route index element={<Quession/>} />
                </Route>
            </Route>
        )
    )
    return <RouterProvider router={root} />
}