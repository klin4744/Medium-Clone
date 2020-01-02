import React from 'react';
import './Article.css';

export default function Article(props) {
   return (
      <li className='Article'>
         <div className='row'>
            <div className='col-8'>
               <div className='container'>
                  <small className='topic-text'>{props.Topic}</small>
                  <h3 onClick={props.handleClick}>{props.Title}</h3>
                  <p className='content-text' onClick={props.handleClick}>
                     {props.Content.substring(0, 300) + '...'}
                  </p>
                  <br></br>
                  <p className='author-text'>
                     {props.Authorname} in {props.Orgname}
                  </p>
                  <p className='d-inline-block text-left'>
                     {props.Dateposted +
                        ' • ' +
                        (props.Content.length % 22) +
                        ' min read'}
                  </p>
               </div>
            </div>
            <div className='col-4'>
               <img src={props.Articleimgurl} alt='img'></img>
            </div>
         </div>
      </li>
   );
}
