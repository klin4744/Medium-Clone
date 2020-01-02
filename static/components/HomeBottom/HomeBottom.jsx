import React from 'react';
import ArticleList from './ArticleList/ArticleList';
import './HomeBottom.css';
import PopularOnMedium from './PopularOnMedium/PopularOnMedium';

export default function HomeBottom(props) {
   const handleClick = id => {
      props.history.push(`/articles/${id}`);
   };
   return (
      <div className='row HomeBottom'>
         <div className='col-7'>
            <ArticleList handleClick={handleClick} articles={props.articles} />
         </div>
         <div className='col-5'>
            <PopularOnMedium handleClick={handleClick} />
         </div>
      </div>
   );
}
