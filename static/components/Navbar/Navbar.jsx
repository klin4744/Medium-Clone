import React from 'react';
import './Navbar.css';
import BottomHalf from './BottomHalf/BottomHalf';
import { Link } from 'react-router-dom';

export default function Navbar() {
   return (
      <div className='w-100'>
         <nav className='Navbar navbar bg-white navbar-light navbar-expand-lg w-75 mx-auto'>
            <Link to='/' className='navbar-brand'>
               Medium
            </Link>
            <ul className='navbar-nav ml-auto row'>
               <li className='nav-item col-3'>
                  <i className='fa fa-search mt-2'></i>
               </li>
               <li className='nav-item col-3'>
                  <i className='fa fa-bell-o mt-2'></i>
               </li>
               <li className='nav-item col-3'>
                  <a className='btn btn-outline-secondary '>Upgrade</a>
               </li>
            </ul>
         </nav>
         <BottomHalf />
      </div>
   );
}
