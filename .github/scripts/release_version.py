"""Choose a semantic release tag from Git history."""
import os
import re
import subprocess


def git(*args):
    return subprocess.check_output(['git', *args], text=True).strip()


def next_version(tags, bump):
    versions = [tuple(map(int, match.groups())) for tag in tags
                if (match := re.fullmatch(r'v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)', tag))]
    if not versions:
        return 'v0.0.1'
    major, minor, patch = max(versions)
    if bump == 'major':
        major, minor, patch = major + 1, 0, 0
    elif bump == 'minor':
        minor, patch = minor + 1, 0
    elif bump == 'patch':
        patch += 1
    else:
        raise ValueError(f'Invalid version bump: {bump}')
    return f'v{major}.{minor}.{patch}'


def main():
    tags = git('tag', '--list', 'v*').splitlines()
    semantic_tags = [tag for tag in tags if re.fullmatch(r'v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)', tag)]
    current_tags = set(git('tag', '--points-at', 'HEAD').splitlines())
    existing = [tag for tag in semantic_tags if tag in current_tags]
    if existing:
        tag = max(existing, key=lambda value: tuple(map(int, value[1:].split('.'))))
    else:
        latest = max(semantic_tags, key=lambda value: tuple(map(int, value[1:].split('.'))), default=None)
        history = git('log', '--format=%B', f'{latest}..HEAD' if latest else 'HEAD')
        markers = re.findall(r'^Release-Bump:\s*(major|minor)\s*$', history, re.MULTILINE | re.IGNORECASE)
        bump = os.environ.get('VERSION_BUMP', 'patch')
        if 'major' in [marker.lower() for marker in markers]:
            bump = 'major'
        elif bump == 'patch' and 'minor' in [marker.lower() for marker in markers]:
            bump = 'minor'
        tag = next_version(tags, bump)
    with open(os.environ['GITHUB_OUTPUT'], 'a') as output:
        output.write(f'tag={tag}\n')
    print(f'Release version: {tag}')


if __name__ == '__main__':
    main()
