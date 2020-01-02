import React, { useState, useEffect } from 'react';
import './SingleArticle.css';
import axios from 'axios';

export default function SingleArticle() {
   const [article, changeArticle] = useState(null);
   useEffect(() => {
      if (!article) {
         async function getArticle() {
            const id = window.location.pathname.split('/')[2];
            const { data: article } = await axios.get(`/api/articles/${id}`);
            changeArticle(article[0]);
         }
         getArticle();
      }
   }, [article]);
   return (
      <div className='SingleArticle w-50 mx-auto p-4'>
         {article ? (
            <>
               <h1>{article.Title}</h1>
               <div className='row my-2'>
                  <div className='col-9 mr-auto row'>
                     <div className='col-1'>
                        <img
                           className='author-img mx-auto'
                           src={article.Authorimgurl}
                           alt='author img'
                        ></img>
                     </div>
                     <div className='col-9 pl-4 pt-1'>
                        {article.Authorname} <br />
                        <span className='date'>
                           {article.Dateposted +
                              ' • ' +
                              (article.Content.length % 22) +
                              ' min read'}
                        </span>
                     </div>
                  </div>
                  <div className='col-3 ml-auto'>
                     <i className='fa fa-twitter'></i>
                     <i className='fa fa-fb'></i>
                  </div>
               </div>
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
