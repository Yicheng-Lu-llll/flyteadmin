git add .
git commit -s -m "develop"
git push


current_tag=$(git describe --tags `git rev-list --tags --max-count=1`)

git tag -d $current_tag
git push origin :refs/tags/$current_tag


new_tag=current_tag + 0.01
echo !!!!!!!!!!!!
echo $new_tag
echo !!!!!!!!!!!!
git tag $new_tag
git push origin $new_tag
