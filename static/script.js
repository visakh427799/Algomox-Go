const currentPage=()=>{
                  const queryString = window.location.search;
                  console.log(queryString);
                  const urlParams = new URLSearchParams(queryString);
                  let code=urlParams.get('code')
                  let filename=urlParams.get('filename')
                  let page=urlParams.get('pagenum')
                  return {code,filename,page}
                }

    // const prev=()=>{

    // }
    const pg=(pagenumber)=>{
        var obj=currentPage()
        var code=obj.code;
        var filename=obj.filename;
        
      location.href=`/readAirport?code=${{code}}&filename=${{filename}}&pagenum=${{pagenumber}}`
        
      }

    //   const next=()=>{
        
    //   }


               

            