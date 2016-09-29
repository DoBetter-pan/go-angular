/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

admin.controller('NewBlogCtrl', ['$scope', '$location', 'BlogSrv', 'sections', 'categories', function($scope, $location, BlogSrv, sections, categories){
    $scope.sections = sections;
    $scope.categories = categories;
    $scope.categoriesBySec = categories.slice(0);
    $scope.changeSeciton = function(secId) {
        $scope.categoriesBySec.splice(0, $scope.categoriesBySec.length);
        $scope.categories.forEach(function(e){
            if(e.sectionId == secId){
                $scope.categoriesBySec.push(e);
            }
        });

    };
    $scope.blog = new BlogSrv({
        id: -1
    });

    $scope.save = function(){
        $scope.blog.$save(function(blog){
             $location.path('/view/' + blog.id);
        });
    };    
    $scope.$on('$viewContentLoaded', function(){
        var ng_writer = new Simditor({ 
            textarea: $('#content'),
            placeholder: "",
            toolbar: [
                'title',
                'bold',
                'italic',
                'underline',
                'strikethrough',
                'fontScale',
                'color',
                'ol',             
                'ul',            
                'blockquote',
                'code',         
                'table',
                'link',
                'image',
                'hr',          
                'indent',
                'outdent',
                'alignment'
            ],
            toolbarFloat: true,
            toolbarFloatOffset: 0,
            toolbarHidden: false,
            defalutImage: "images/image.png, images/image.jpg",
            tabIndent: true,
            upload: {
                url: '',
                params: null,
                fileKey: 'upload_file',
                connectionCount: 3,
                leaveConfirm: '正在上传文件, 你确定要离开当前页面么?'
            },
            pasteImage: true,
            cleanPaste: false,
            imageButton: [
                'upload',
                'external'
            ],
            allowedTags: [
                'br',
                'span',
                'a',
                'img',
                'b',
                'strong',
                'i',
                'strike',
                'u',
                'font',
                'p',
                'ul',
                'ol',
                'li',
                'blockquote',
                'pre',
                'code',
                'h1',
                'h2',
                'h3',
                'h4',
                'hr'
            ],
            allowedAttributes: {
                img: ['src', 'alt', 'width', 'height', 'data-non-image'],
                a: ['href', 'target'],
                font: ['color'],
                code: ['class']
            },
            allowedStyles: {
                span: ['color', 'font-size'],
                b: ['color'],
                i: ['color'],
                strong: ['color'],
                strike: ['color'],
                u: ['color'],
                p: ['margin-left', 'text-align'],
                h1: ['margin-left', 'text-align'],
                h2: ['margin-left', 'text-align'],
                h3: ['margin-left', 'text-align'],
                h4: ['margin-left', 'text-align']
            },
            codeLanguages: [
                { name: 'Bash', value: 'bash' },
                { name: 'C++', value: 'c++' },
                { name: 'C#', value: 'cs' },
                { name: 'CSS', value: 'css' },
                { name: 'Erlang', value: 'erlang' },
                { name: 'Less', value: 'less' },
                { name: 'Sass', value: 'sass' },
                { name: 'Diff', value: 'diff' },
                { name: 'CoffeeScript', value: 'coffeescript' },
                { name: 'HTML,XML', value: 'html' },
                { name: 'JSON', value: 'json' },
                { name: 'Java', value: 'java' },
                { name: 'JavaScript', value: 'js' },
                { name: 'Markdown', value: 'markdown' },
                { name: 'Objective C', value: 'oc' },
                { name: 'PHP', value: 'php' },
                { name: 'Perl', value: 'parl' },
                { name: 'Python', value: 'python' },
                { name: 'Ruby', value: 'ruby' },
                { name: 'SQL', value: 'sql'},
            ],
            params: {}
        });
    });
}]);

admin.controller('EditCtrl', ['$scope', '$location', 'article', function($scope, $location, article){
    $scope.article = article;

    $scope.save = function(){
        $scope.article.$save(function(article){
            $location.path('/view/' + article.id);
        });
    };

    $scope.remove = function(){
        $scope.article.$remove(function(article){
            $location.path('/');
        });
    };
}]);

admin.controller('NewCtrl', ['$scope', '$location', 'BlogSrv', function($scope, $location, BlogSrv){
    $scope.article = new BlogSrv({
        id: -1
    });

    $scope.save = function(){
        $scope.article.$save(function(article){
            $location.path('/view/' + article.id);
        });
    };
}]);