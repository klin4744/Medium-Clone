import React, { useState, useEffect } from 'react';
import PopContainer from './PopContainer/PopContainer';
import axios from 'axios';
import './PopularOnMedium.css';
import BottomPop from './BottomPop/BottomPop';

function stickyBottom(component, height) {
   if (window.pageYOffset >= height) {
      component.classList.add('sticky-2');
   } else {
      component.classList.remove('sticky-2');
   }
}

function handleScroll() {
   const component = document.getElementById('PopularOnMedium');
   stickyBottom(component, 570);
}

export default function PopularOnMedium() {
   const [topArticles, setTop] = useState([]);
   useEffect(() => {
      if (!topArticles.length) {
         async function getTopArticles() {
            const { data: articles } = await axios.get('/api/articles/top');
            setTop(articles);
         }
         getTopArticles();
      }
   }, [topArticles]);
   useEffect(() => {
      window.addEventListener('scroll', handleScroll);

      return () => {
         window.removeEventListener('scroll', handleScroll);
      };
   }, [handleScroll]);

   return (
      <div className='container PopularOnMedium' id='PopularOnMedium'>
         <h3 id='title-text'>Popular On Medium</h3>
         {topArticles.map((article, idx) => (
            <PopContainer {...article} idx={idx} />
         ))}
         <BottomPop />
      </div>
   );
}
