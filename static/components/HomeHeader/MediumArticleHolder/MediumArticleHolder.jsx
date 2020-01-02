import React from 'react';
import './MediumArticleHolder.css';

export default function MediumArticleHolder(props) {
   return (
      <div className={`col-${props.size}  MediumArticleHolder`}>
         <img src={props.imgUrl} />
         <div className='card-content'>
            <h3 className='content-text' onClick={props.handleClick}>
               {props.title}
            </h3>
            <p className='content-text' onClick={props.handleClick}>
               {props.content.slice(0, 200) + '...'}
            </p>
            <small>
               {props.author} in {props.location}
            </small>
            <p>{props.date + ' • ' + props.time + ' min read'}</p>
         </div>
      </div>
   );
}
