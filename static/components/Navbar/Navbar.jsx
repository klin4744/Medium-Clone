import React from 'react';
import './Navbar.css';
import BottomHalf from './BottomHalf/BottomHalf';

export default function Navbar() {
   return (
      <div className='w-75 mx-auto'>
         <nav className='Navbar navbar bg-white navbar-light navbar-expand-lg'>
            <a className='navbar-brand'>Medium</a>
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
