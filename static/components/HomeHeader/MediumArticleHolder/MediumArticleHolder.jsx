import React from 'react';
import './MediumArticleHolder.css';

export default function MediumArticleHolder(props) {
   return (
      <div className={`col-${props.size}  MediumArticleHolder`}>
         <img src={props.imgSrc} />
         <div className='card-content'>
            <h3>{props.title}</h3>
            <p>{props.content.slice(0, 200) + '...'}</p>
            <small>
               {props.author} in {props.location}
            </small>
            <p>{props.date + ' • ' + props.time + 'min read'}</p>
         </div>
      </div>
   );
}
