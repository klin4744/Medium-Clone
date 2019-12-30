import React, { useState, useEffect } from 'react';
import PopContainer from './PopContainer/PopContainer';
import axios from 'axios';
import './PopularOnMedium.css';

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
   }, topArticles);

   return (
      <div className='container PopularOnMedium'>
         <h3 id='title-text'>Popular On Medium</h3>
         {topArticles.map((article, idx) => (
            <PopContainer {...article} idx={idx} />
         ))}
      </div>
   );
}
