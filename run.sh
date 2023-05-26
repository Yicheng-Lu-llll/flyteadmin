git add .
git commit -s -m "develop"
git push


current_tag=$(git describe --tags `git rev-list --tags --max-count=1`)

git tag -d $current_tag
git push origin :refs/tags/$current_tag


new_tag=$(echo $latest_tag | awk -F. '{$NF = $NF + 1;} 1' | sed 's/ /./g')
echo $new_tag
git tag $new_tag
git push origin $new_tag
