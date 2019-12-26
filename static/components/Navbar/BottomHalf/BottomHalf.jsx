import React from 'react';
import Slider from 'react-slick';
import ContentTab from './ContentTab/ContentTab';
import './BottomHalf.css';

function stickyBottom(component, height) {
   if (window.pageYOffset >= height) {
      component.classList.add('sticky');
      component.classList.add('bg-white');
      component.classList.add('border-bottom');
   } else {
      component.classList.remove('sticky');
      component.classList.remove('bg-white');
      component.classList.remove('border-bottom');
   }
}

class BottomHalf extends React.Component {
   componentDidMount() {
      const component = document.getElementById('BottomHalf');
      window.addEventListener('scroll', () => stickyBottom(component, 75));
   }
   componentWillUnmount() {
      window.removeEventListener(stickyBottom);
   }
   render() {
      let settings = {
         dots: false,
         infinite: false,
         speed: 500,
         slidesToShow: 12,
         slidesToScroll: 1,
      };
      return (
         <div className='BottomHalf' id='BottomHalf'>
            <div className='container'>
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
         </div>
      );
   }
}

export default BottomHalf;
