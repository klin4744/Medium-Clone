import React from 'react';
import ArticleList from './ArticleList/ArticleList';
import './HomeBottom.css';

// const articles = [
//    {
//       title: 'React JS for Beginners — The Basics',
//       topic: 'JavaScript',
//       author: 'Cem Eygi',
//       content:
//          "React JS is today's most popular JavaScript Library for building User Interfaces, which has created by Facebook. We can build modern, fast Single Page Applications or websites with React.",
//       imgUrl: 'https://miro.medium.com/max/2400/1*y6C4nSvy2Woe0m7bWEn4BA.png',
//       date: 'March 25, 2019',
//       time: '5',
//       location: 'codeburst.io',
//    },
//    {
//       title: 'React JS for Beginners — The Basics',
//       topic: 'JavaScript',
//       author: 'Cem Eygi',
//       content:
//          "React JS is today's most popular JavaScript Library for building User Interfaces, which has created by Facebook. We can build modern, fast Single Page Applications or websites with React.",
//       imgUrl: 'https://miro.medium.com/max/2400/1*y6C4nSvy2Woe0m7bWEn4BA.png',
//       date: 'March 25, 2019',
//       time: '5',
//       location: 'codeburst.io',
//    },
//    {
//       title: 'React JS for Beginners — The Basics',
//       topic: 'JavaScript',
//       author: 'Cem Eygi',
//       content:
//          "React JS is today's most popular JavaScript Library for building User Interfaces, which has created by Facebook. We can build modern, fast Single Page Applications or websites with React.",
//       imgUrl: 'https://miro.medium.com/max/2400/1*y6C4nSvy2Woe0m7bWEn4BA.png',
//       date: 'March 25, 2019',
//       time: '5',
//       location: 'codeburst.io',
//    },
//    {
//       title: 'React JS for Beginners — The Basics',
//       topic: 'JavaScript',
//       author: 'Cem Eygi',
//       content:
//          "React JS is today's most popular JavaScript Library for building User Interfaces, which has created by Facebook. We can build modern, fast Single Page Applications or websites with React.",
//       imgUrl: 'https://miro.medium.com/max/2400/1*y6C4nSvy2Woe0m7bWEn4BA.png',
//       date: 'March 25, 2019',
//       time: '5',
//       location: 'codeburst.io',
//    },
//    {
//       title: 'React JS for Beginners — The Basics',
//       topic: 'JavaScript',
//       author: 'Cem Eygi',
//       content:
//          "React JS is today's most popular JavaScript Library for building User Interfaces, which has created by Facebook. We can build modern, fast Single Page Applications or websites with React.",
//       imgUrl: 'https://miro.medium.com/max/2400/1*y6C4nSvy2Woe0m7bWEn4BA.png',
//       date: 'March 25, 2019',
//       time: '5',
//       location: 'codeburst.io',
//    },
// ];

export default function HomeBottom(props) {
   return (
      <div className='row HomeBottom'>
         <div className='col-7'>
            <ArticleList articles={props.articles} />
         </div>
         <div className='col-5'>
            <PopularOnMedium />
         </div>
      </div>
   );
}
