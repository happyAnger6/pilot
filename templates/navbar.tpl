{{define "navbar"}}
 <div class="container">

      <!-- Static navbar -->
      <nav class="navbar navbar-default">
        <div class="container-fluid">
          <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
              <span class="sr-only">Toggle navigation</span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">Pilot容器管理平台alpha版</a>
          </div>
          <div id="navbar" class="navbar-collapse collapse">
            <ul class="nav navbar-nav">
              <li><a href="/">主 页</a></li>
              <li><a href="/list/boards">板　子</a></li>
              <li><a href="/hosts">主 机</a></li>
              <li class="dropdown">
                <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">部署simware <span class="caret"></span></a>
                <ul class="dropdown-menu">
                  <li><a href="/deploy/createTemplate">创建模板</a></li>
                  <li><a href="/deploy/startBoard">启动板子</a></li>
                  <li role="separator" class="divider"></li>
                  <li class="dropdown-header">组网相关</li>
                  <li><a href="#">连线</a></li>
                </ul>
              </li>
            </ul>
          </div><!--/.nav-collapse -->
        </div><!--/.container-fluid -->
      </nav>

    </div>
    {{end}}