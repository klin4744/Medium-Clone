import React from 'react';
import './Article.css';

export default function Article(props) {
   console.log(props);
   return (
      <li className='Article'>
         <div className='row'>
            <div className='col-8'>
               <div className='container'>
                  <small className='topic-text'>{props.topic}</small>
                  <h3>{props.title}</h3>
                  <p>{props.content.substring(0, 300) + '...'}</p>
                  <br></br>
                  <p className='author-text'>
                     {props.author} in {props.location}
                  </p>
                  <p className='d-inline-block text-left'>
                     {props.date + ' • ' + props.time + ' min read'}
                  </p>
               </div>
            </div>
            <div className='col-4'>
               <img src={props.imgUrl} alt='img'></img>
            </div>
         </div>
      </li>
   );
}
