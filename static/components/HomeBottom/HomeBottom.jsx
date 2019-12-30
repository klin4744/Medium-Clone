import React from 'react';
import ArticleList from './ArticleList/ArticleList';
import './HomeBottom.css';
import PopularOnMedium from './PopularOnMedium/PopularOnMedium';

export default function HomeBottom(props) {
   return (
      <div className='row HomeBottom'>
         <div className='col-7'>
            <ArticleList articles={props.articles} />
         </div>
         <div className='col-5'>
            <PopularOnMedium />
         </div>
      </div>
   );
}
