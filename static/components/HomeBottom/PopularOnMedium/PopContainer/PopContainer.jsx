import React from 'react';
import './PopContainer.css';

export default function PopContainer(props) {
   return (
      <div className='row PopContainer w-75'>
         <div className='col-2'>
            <h3 className='place-text'>{'0' + (props.idx + 1)}</h3>
         </div>
         <div className='col-10'>
            <h3 className='content-text' onClick={props.handleClick}>
               {props.Title}
            </h3>
            <p className='author-text'>
               {props.Authorname} in {props.Orgname}
            </p>
            <p className='d-inline-block text-left time-text'>
               {props.Dateposted +
                  ' • ' +
                  (props.Content.length % 22) +
                  ' min read'}
            </p>
         </div>
      </div>
   );
}
