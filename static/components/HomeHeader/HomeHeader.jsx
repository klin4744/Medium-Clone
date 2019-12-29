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
            { this.state.articles.length ? <MediumArticleHolder
               size='4'
               imgUrl={this.state.articles[0].Articleimgurl}
               title={this.state.articles[0].Title}
               content={this.state.articles[0].Content}
               date={this.state.articles[0].Dateposted}
               time={this.state.articles[0].Content.length % 20}
               author={this.state.articles[0].Authorname}
               location={this.state.articles[0].Orgname}
            /> : <></>}
            <Display articles={this.state.articles} />
            {this.state.articles.length ? <MediumArticleHolder
               size='4'
               imgUrl={this.state.articles[2].Articleimgurl}
               title={this.state.articles[2].Title}
               content={this.state.articles[2].Content}
               date={this.state.articles[2].Dateposted}
               time={this.state.articles[2].Content.length % 20}
               author={this.state.articles[2].Authorname}
               location={this.state.articles[2].Orgname}
            /> : <></>}
            <a id='editor'>See Editor's Picks ></a>
            <div className='border-bottom'></div>
         </div>
      );
   }
}
export default HomeHeader;
