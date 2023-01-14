<p>docker build -t transferauth .</p>
<p>docker run -d -p 9091:9091  -p 8888:8888 transferauth  </p>


<p>docker run --name some-postgres -p 54322:5432 -e POSTGRES_PASSWORD=pgpwd4habr -e POSTGRES_DB=habrdb -e POSTGRES_USER=habrpguser -d postgres 
</p>