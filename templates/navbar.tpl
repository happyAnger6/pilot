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
            <a class="navbar-brand" href="/">Cloudware 0.1</a>
          </div>
          <div id="navbar" class="navbar-collapse collapse">
            <ul class="nav navbar-nav">
              <!-- <li><a href="/list/boards"></a></li>
              <li><a href="/hosts">主 机</a></li> -->
              <li class="dropdown">
                <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">simware <span class="caret"></span></a>
                <ul class="dropdown-menu">
                  <!-- <li>创建模板</li> -->
                  <li><a href="/deploy/startBoard">部署</a></li>
                  <li role="separator" class="divider"></li>
                  <li class="dropdown-header">查看</li>
                  <li><a href="/list/boards">按板查看</a></li>
                  <li><a href="/list/devices">按设备查看</a></li>
                  <li><a href="/list/network/connections">组网信息</a></li>
                </ul>
              </li>
            </ul>
            <ul class="nav navbar-nav navbar-right">
            <li><a>你　好 !  {{.UserName}}</a></li>
            <li><a href="/logout">退出登陆</a></li>
            </ul>
          </div><!--/.nav-collapse -->
        </div><!--/.container-fluid -->
      </nav>

    </div>
    {{end}}