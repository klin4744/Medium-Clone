import React from 'react';
import Slider from 'react-slick';
import ContentTab from './ContentTab/ContentTab';
import './BottomHalf.css';

export default function BottomHalf() {
   let settings = {
      dots: false,
      infinite: false,
      speed: 500,
      slidesToShow: 12,
      slidesToScroll: 1,
   };
   return (
      <div className='container mx-auto BottomHalf'>
         <Slider {...settings} className='text-center'>
            <ContentTab content='Home' />
            <ContentTab content='Gift Medium' />
            <ContentTab content='One Zero' />
            <ContentTab content='Elemental' />
            <ContentTab content='Gen' />
            <ContentTab content='Zora' />
            <ContentTab content='Force' />
            <ContentTab content='Human Parts' />
            <ContentTab content='Marker' />
            <ContentTab content='Level' />
            <ContentTab content='Heated' />
            <ContentTab content='Modus' />
            <ContentTab content='More' />
         </Slider>
      </div>
   );
}
