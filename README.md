# test_auto_push


update plan:
* auto clean .git cache. In each epoch(500), contains 500 times commit, it would auto delete the commit files and relation tree files, only remaining 1 version commit.
> formulation is: `del(iter(flag) - set(cur+last))`
>   1. calculate `set(cur+last)`, cur is the lastest version commit, last is the version would be remain(cur version-500)
>   2. flag is a pointer , point each line of version list which would be deleted
>   3. calculate `iter(flag) - set(cur+last)`, the difference of each inner version with `set(cur+last)`
>   4. delete the difference of each redundent version
