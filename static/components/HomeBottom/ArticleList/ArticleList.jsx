import React from 'react';
import Article from './Article/Article';
import './ArticleList.css';

export default function ArticleList(props) {
   return (
      <ul className='ArticleList'>
         {props.articles.map(article => (
            <Article {...article} />
         ))}
      </ul>
   );
}
