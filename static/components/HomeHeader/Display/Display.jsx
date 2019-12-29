import React from 'react';
import SmallContentContainer from './SmallContentContainer/SmallContentContainer.jsx';

export default function Display(props) {
   return (
      <div className='col-4'>
         {props.articles.map(article => (
            <SmallContentContainer  imgUrl={article.Articleimgurl} content={article.Content} date={article.Dateposted} title={article.Title} time={article.Content.length % 20} author={article.Authorname} location={article.Orgname} />
         ))}
      </div>
   );
}
