import React from 'react';
import MediumArticleHolder from './MediumArticleHolder/MediumArticleHolder';
import Display from './Display/Display.jsx';
import './HomeHeader.css';
import axios from 'axios';
import Article from '../HomeBottom/ArticleList/Article/Article';

class HomeHeader extends React.Component {
   constructor(props) {
      super(props);
      this.state = {
         articles: [],
      };
   }
   async componentDidMount() {
      const { data: articles } = await axios.get('/api/articles');
      this.setState({ articles });

   }
   render() {
      return (
         <div className='row HomeHeader'>
            <MediumArticleHolder
               size='4'
               imgSrc='https://miro.medium.com/max/6528/1*Kd0UGiDvgooFooCy28rs8Q.jpeg'
               title='The Top 10 Things Wrong with JavaScript'
               content='JavaScript has a reputation for being one of the worst programming languages in existence, and for good reasons! JavaScript is easy to learn and easy to use, except when it’s not. There are many “gotchas” that can trip you up. Below, I glean some of the best from various online sources…'
               date='Dec 17, 2018'
               time='12'
               author='Milap Neupane'
               location='freeCodeCamp.org'
            />
            <Display articles={this.state.articles} />
            <MediumArticleHolder
               size='4'
               imgSrc='https://miro.medium.com/max/6528/1*Kd0UGiDvgooFooCy28rs8Q.jpeg'
               title='The Top 10 Things Wrong with JavaScript'
               content='JavaScript has a reputation for being one of the worst programming languages in existence, and for good reasons! JavaScript is easy to learn and easy to use, except when it’s not. There are many “gotchas” that can trip you up. Below, I glean some of the best from various online sources…'
               date='Dec 17, 2018'
               time='12'
               author='Milap Neupane'
               location='freeCodeCamp.org'
            />
            <a id='editor'>See Editor's Picks ></a>
            <div className='border-bottom'></div>
         </div>
      );
   }
}
export default HomeHeader;
