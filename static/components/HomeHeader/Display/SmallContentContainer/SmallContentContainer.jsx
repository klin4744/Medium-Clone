import React from 'react';
import "./SmallContentContainer.css"

export default function SmallContentContainer(props) {
   return (
      <div className='row SmallContentContainer'>
         <div className='col-3'>
            <img src={props.imgUrl} alt='content image' />
         </div>
         <div className='col-8 container'>
            <h3>{props.title}</h3>
            <p>{props.content.slice(0, 20) + '...'}</p>
            <small>
               {props.author} in {props.location}
            </small>
            <p>{props.date + ' • ' + props.time + 'min read'}</p>
         </div>
      </div>
   );
};
