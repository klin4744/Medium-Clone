import React from 'react';
import MediumArticleHolder from './MediumArticleHolder/MediumArticleHolder';
import Display from './Display/Display.jsx';
import './HomeHeader.css';

function HomeHeader(props) {
   return (
      <div className='row HomeHeader'>
         {props.articles.length ? (
            <MediumArticleHolder
               size='4'
               imgUrl={props.articles[0].Articleimgurl}
               title={props.articles[0].Title}
               content={props.articles[0].Content}
               date={props.articles[0].Dateposted}
               time={props.articles[0].Content.length % 20}
               author={props.articles[0].Authorname}
               location={props.articles[0].Orgname}
            />
         ) : (
            <></>
         )}
         <Display articles={props.articles} />
         {props.articles.length ? (
            <MediumArticleHolder
               size='4'
               imgUrl={props.articles[2].Articleimgurl}
               title={props.articles[2].Title}
               content={props.articles[2].Content}
               date={props.articles[2].Dateposted}
               time={props.articles[2].Content.length % 20}
               author={props.articles[2].Authorname}
               location={props.articles[2].Orgname}
            />
         ) : (
            <></>
         )}
         <a id='editor'>See Editor's Picks ></a>
         <div className='border-bottom'></div>
      </div>
   );
}
export default HomeHeader;
