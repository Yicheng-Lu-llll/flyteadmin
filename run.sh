git add .
git commit -s -m "develop"
git push

git tag -d v1.0.0
git push origin :refs/tags/v1.0.1

git tag v1.0.0
git push origin v1.0.0
