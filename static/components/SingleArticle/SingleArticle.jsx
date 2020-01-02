import React, { useState, useEffect } from 'react';
import './SingleArticle.css';
import axios from 'axios';

export default function SingleArticle() {
   const [article, changeArticle] = useState(null);
   useEffect(() => {
      if (!article) {
         async function getArticle() {
            const id = window.location.pathname.split('/')[1];
            const { data: article } = await axios.get(`/api/articles/${id}`);
            changeArticle(article);
         }
         getArticle();
      }
   }, [article]);
   return (
      <div className='SingleArticle w-75 mx-auto p-4'>
         {article ? (
            <>
               <h1>{article.Title}</h1>
               <img className='article-img' src={article.Articleimgurl}></img>
               <p>{article.Content}</p>
            </>
         ) : (
            <></>
         )}
      </div>
   );
}

// size='4'
// imgUrl={props.articles[0].Articleimgurl}
// title={props.articles[0].Title}
// content={props.articles[0].Content}
// date={props.articles[0].Dateposted}
// time={props.articles[0].Content.length % 20}
// author={props.articles[0].Authorname}
// location={props.articles[0].Orgname}
